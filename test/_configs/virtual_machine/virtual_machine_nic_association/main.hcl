config {
  location = "ireland"
  clouds   = ["aws", "azure"]
}
multy "virtual_network" "example_vn" {
  name       = "example_vn"
  cidr_block = "10.0.0.0/16"
}
multy "subnet" "subnet" {
  name               = "subnet"
  cidr_block         = "10.0.2.0/24"
  virtual_network = example_vn
  availability_zone  = 2
}
multy "network_interface" "nic" {
  name      = "test-nic"
  subnet_id = subnet
}
multy "virtual_machine" "vm" {
  name                  = "test-vm"
  os                    = "linux"
  size                  = "micro"
  user_data             = "echo 'Hello World'"
  subnet_id             = subnet
  network_interface_ids = [nic]
  public_ip             = false
}