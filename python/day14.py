import unittest

from math import ceil
from typing import Any, Dict, Iterable, List, Set, Tuple

from common import fixture


class Component:
    """
    Component describes a substance and its coefficient.
    """

    def __init__(self, substance: str, coefficient: int = 1) -> None:
        assert substance
        assert coefficient > 0
        self.substance = substance
        self.coefficient = coefficient

    def __str__(self) -> str:
        if self.coefficient == 1:
            return self.substance
        return f"{self.coefficient}{self.substance}"

    @classmethod
    def from_str(cls, s: str) -> "Component":
        coefficient, substance = s.split(" ")
        return Component(substance=substance.strip(), coefficient=int(coefficient))


class Expression:
    """
    Expression describes a list of reactant or product components.
    """

    def __init__(self, components: Iterable[Component]) -> None:
        self.components = [
            component for component in components if component.coefficient
        ]

    def __contains__(self, substance: str) -> bool:
        return any(component.substance == substance for component in self.components)

    def __getitem__(self, substance: str) -> Component:
        for component in self.components:
            if component.substance == substance:
                return component
        raise KeyError(substance)

    def __mul__(self, n: Any) -> "Expression":
        if not isinstance(n, int):
            return NotImplemented
        return Expression(
            components=[
                Component(component.substance, component.coefficient * n)
                for component in self.components
            ]
        )

    def __str__(self) -> str:
        return " + ".join(str(component) for component in self.components)

    @classmethod
    def from_str(cls, s: str) -> "Expression":
        components = s.split(", ")
        return Expression(
            components=[Component.from_str(component) for component in components]
        )


class Reaction:
    """
    Reaction describes the reactants and products of a possible reaction.
    """

    def __init__(self, reactants: Expression, products: Expression) -> None:
        self.reactants = reactants
        self.products = products

    def __mul__(self, n: Any) -> "Reaction":
        if not isinstance(n, int):
            return NotImplemented
        return Reaction(
            reactants=self.reactants * n,
            products=self.products * n,
        )

    def __str__(self) -> str:
        return f"{self.reactants} → {self.products}"

    @property
    def is_composition(self) -> bool:
        return len(self.products.components) == 1

    @classmethod
    def from_str(cls, s: str) -> "Reaction":
        reactants, products = s.split(" => ")
        return Reaction(
            reactants=Expression.from_str(reactants),
            products=Expression.from_str(products),
        )


class CompositionTable:
    """
    Container for reactions that produce a single substance, indexed by the product.
    """

    def __init__(self, reactions: Iterable[Reaction]) -> None:
        for reaction in reactions:
            assert reaction.is_composition
        self.reactions = {
            reaction.products.components[0].substance: reaction
            for reaction in reactions
        }

    def __contains__(self, product: str) -> bool:
        return product in self.reactions

    def __getitem__(self, product: str) -> Reaction:
        return self.reactions[product]

    def __str__(self) -> str:
        return "\n".join([str(reaction) for reaction in self.reactions.values()])

    @classmethod
    def from_str(cls, s: str) -> "CompositionTable":
        return CompositionTable(
            reactions=[Reaction.from_str(line) for line in s.splitlines()]
        )

    @classmethod
    def from_file(cls, name: str) -> "CompositionTable":
        with open(name, "r") as fp:
            return CompositionTable.from_str(fp.read())


def incr(v: Dict[str, int], key: str, n: int) -> None:
    if key in v:
        v[key] += n
    else:
        v[key] = n


def decr(v: Dict[str, int], key: str, n: int) -> None:
    incr(v, key, -n)


def expand(reaction: Reaction, compositions: CompositionTable) -> Reaction:
    """
    Expand a reaction as much as possible given the synthesis reactions for its reactants.
    """
    final_reactants: Dict[str, int] = {}
    final_products: Dict[str, int] = {
        component.substance: component.coefficient
        for component in reaction.products.components
    }

    # queue all reactants to be expanded
    reactants: Dict[str, int] = {
        component.substance: component.coefficient
        for component in reaction.reactants.components
    }
    while sum(reactants.values()):
        for substance, coefficient in reactants.items():
            if coefficient == 0:
                # we don't need any more of this substance
                continue

            if substance not in compositions:
                # we cannot expand this reactant any further
                incr(final_reactants, substance, coefficient)
                decr(reactants, substance, coefficient)
                continue

            # how many batches do we need of this reactant recipe?
            formula = compositions[substance]
            batch_size = formula.products.components[0].coefficient
            batch_count = ceil(coefficient / batch_size)

            # enqueue subcomponents to reactant list
            for reactant in formula.reactants.components:
                needed = reactant.coefficient * batch_count
                if reactant.substance in final_products:
                    # steal from previous byproducts
                    available = final_products[reactant.substance]
                    taken = min(needed, available)
                    needed -= taken
                    decr(final_products, reactant.substance, taken)
                incr(reactants, reactant.substance, needed)

            # add excess from the batch to the product list
            output = batch_size * batch_count
            excess = output - coefficient
            incr(final_products, substance, excess)

            # remove the decomposed reactant from reactant list
            decr(reactants, substance, output - excess)
            assert reactants[substance] == 0

            # start again since dict may have been modified
            break

    # return expanded reaction
    return Reaction(
        reactants=Expression(
            components=[
                Component(substance=substance, coefficient=coefficient)
                for substance, coefficient in final_reactants.items()
                if coefficient
            ]
        ),
        products=Expression(
            components=[
                Component(substance=substance, coefficient=coefficient)
                for substance, coefficient in final_products.items()
                if coefficient
            ]
        ),
    )


def compute_ore_consumed(compositions: CompositionTable, fuel: int = 1) -> int:
    """
    Compute ORE required to produce 1x FUEL.
    """
    reaction = compositions["FUEL"] * fuel
    reaction = expand(reaction, compositions)
    for reactant in reaction.reactants.components:
        if reactant.substance == "ORE":
            return reactant.coefficient
    raise RuntimeError("No ORE?")


def compute_fuel_yield(compositions: CompositionTable, max_ore: int) -> int:
    """
    Compute the maximum FUEL yield given max_ore.

    Binary search from [min_fuel, min_fuel * 2)
    """
    ore_per_fuel = compute_ore_consumed(compositions)
    min_fuel = int(max_ore / ore_per_fuel)
    max_fuel = min_fuel * 2
    while min_fuel < max_fuel - 1:
        n = min_fuel + int((max_fuel - min_fuel) / 2)
        ore = compute_ore_consumed(compositions, fuel=n)
        if ore == max_ore:
            return n
        if ore < max_ore:
            min_fuel = n
        if ore > max_ore:
            max_fuel = n
    return min_fuel


class TestDay14(unittest.TestCase):
    def test_part1_example1(self):
        compositions = CompositionTable.from_str(
            "10 ORE => 10 A\n"
            "1 ORE => 1 B\n"
            "7 A, 1 B => 1 C\n"
            "7 A, 1 C => 1 D\n"
            "7 A, 1 D => 1 E\n"
            "7 A, 1 E => 1 FUEL\n"
        )
        self.assertEqual(compute_ore_consumed(compositions), 31)

    def test_part1_example2(self):
        compositions = CompositionTable.from_str(
            "9 ORE => 2 A\n"
            "8 ORE => 3 B\n"
            "7 ORE => 5 C\n"
            "3 A, 4 B => 1 AB\n"
            "5 B, 7 C => 1 BC\n"
            "4 C, 1 A => 1 CA\n"
            "2 AB, 3 BC, 4 CA => 1 FUEL\n"
        )
        self.assertEqual(compute_ore_consumed(compositions), 165)

    def test_part1_example3(self):
        compositions = CompositionTable.from_str(
            "157 ORE => 5 NZVS\n"
            "165 ORE => 6 DCFZ\n"
            "44 XJWVT, 5 KHKGT, 1 QDVJ, 29 NZVS, 9 GPVTF, 48 HKGWZ => 1 FUEL\n"
            "12 HKGWZ, 1 GPVTF, 8 PSHF => 9 QDVJ\n"
            "179 ORE => 7 PSHF\n"
            "177 ORE => 5 HKGWZ\n"
            "7 DCFZ, 7 PSHF => 2 XJWVT\n"
            "165 ORE => 2 GPVTF\n"
            "3 DCFZ, 7 NZVS, 5 HKGWZ, 10 PSHF => 8 KHKGT\n"
        )
        self.assertEqual(compute_ore_consumed(compositions), 13312)

    def test_part1_example4(self):
        compositions = CompositionTable.from_str(
            "2 VPVL, 7 FWMGM, 2 CXFTF, 11 MNCFX => 1 STKFG\n"
            "17 NVRVD, 3 JNWZP => 8 VPVL\n"
            "53 STKFG, 6 MNCFX, 46 VJHF, 81 HVMC, 68 CXFTF, 25 GNMV => 1 FUEL\n"
            "22 VJHF, 37 MNCFX => 5 FWMGM\n"
            "139 ORE => 4 NVRVD\n"
            "144 ORE => 7 JNWZP\n"
            "5 MNCFX, 7 RFSQX, 2 FWMGM, 2 VPVL, 19 CXFTF => 3 HVMC\n"
            "5 VJHF, 7 MNCFX, 9 VPVL, 37 CXFTF => 6 GNMV\n"
            "145 ORE => 6 MNCFX\n"
            "1 NVRVD => 8 CXFTF\n"
            "1 VJHF, 6 MNCFX => 4 RFSQX\n"
            "176 ORE => 6 VJHF\n"
        )
        self.assertEqual(compute_ore_consumed(compositions), 180697)

    def test_part1_example5(self):
        compositions = CompositionTable.from_str(
            "171 ORE => 8 CNZTR\n"
            "7 ZLQW, 3 BMBT, 9 XCVML, 26 XMNCP, 1 WPTQ, 2 MZWV, 1 RJRHP => 4 PLWSL\n"
            "114 ORE => 4 BHXH\n"
            "14 VRPVC => 6 BMBT\n"
            "6 BHXH, 18 KTJDG, 12 WPTQ, 7 PLWSL, 31 FHTLT, 37 ZDVW => 1 FUEL\n"
            "6 WPTQ, 2 BMBT, 8 ZLQW, 18 KTJDG, 1 XMNCP, 6 MZWV, 1 RJRHP => 6 FHTLT\n"
            "15 XDBXC, 2 LTCX, 1 VRPVC => 6 ZLQW\n"
            "13 WPTQ, 10 LTCX, 3 RJRHP, 14 XMNCP, 2 MZWV, 1 ZLQW => 1 ZDVW\n"
            "5 BMBT => 4 WPTQ\n"
            "189 ORE => 9 KTJDG\n"
            "1 MZWV, 17 XDBXC, 3 XCVML => 2 XMNCP\n"
            "12 VRPVC, 27 CNZTR => 2 XDBXC\n"
            "15 KTJDG, 12 BHXH => 5 XCVML\n"
            "3 BHXH, 2 VRPVC => 7 MZWV\n"
            "121 ORE => 7 VRPVC\n"
            "7 XCVML => 6 RJRHP\n"
            "5 BHXH, 4 VRPVC => 5 LTCX\n"
        )
        self.assertEqual(compute_ore_consumed(compositions), 2210736)

    def test_part1(self):
        compositions = CompositionTable.from_file(fixture("day14"))
        self.assertEqual(compute_ore_consumed(compositions), 365768)

    def test_part2(self):
        compositions = CompositionTable.from_file(fixture("day14"))
        self.assertEqual(compute_fuel_yield(compositions, 1000000000000), 3756877)
