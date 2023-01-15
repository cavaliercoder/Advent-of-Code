# -*- mode: ruby -*-
# vi: set ft=ruby :

Vagrant.configure("2") do |config|
  config.vm.box = "hashicorp/bionic64"
  config.vm.provision "shell", inline: <<-SHELL
    apt-get update
    apt-get upgrade
    apt-get install -y clang libssl-dev
  SHELL

  config.vm.provider "vmware_desktop" do |v|
    v.vmx["cpuid.coresPerSocket"]  = "1"
    v.vmx["numvcpus"] = "4"
    v.vmx["memsize"] = 4096
  end
end
