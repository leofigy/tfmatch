provider "aws" {
    region = "us-west-2"
}

data "aws_ami" "ubuntu_linux" {
    most_recent=true
    filter {
        name = "name"
        values = ["ubuntu/images/hvm-ssd/ubuntu-trusty-14.04-amd64-server-*"]
    }

    filter {
       name   = "virtualization-type"
       values = ["hvm"]
    }

    owners = ["099720109477"] # Canonical
}

resource "aws_default_subnet" "default_az" {
    availability_zone = "us-west-2a"
    tags = {
        Name ="DefaultSubnet"
    }
}

# Opening port using module http
module "web_server_sg" {
    source = "terraform-aws-modules/security-group/aws//modules/http-80"
    name = "web-server"
    description = "Open port 80"
    vpc_id = "${aws_default_subnet.default_az.vpc_id}"
    ingress_cidr_blocks = ["0.0.0.0/0"]
}

resource "aws_security_group" "admin_sg" {
    name= "leofigy.bot-admin_sg"
    description = "Admin SG"
    vpc_id = "${aws_default_subnet.default_az.vpc_id}"
    ingress{
        from_port = 22
        to_port = 22
        protocol = "tcp"
        cidr_blocks=["0.0.0.0/0"]
    }
}

resource "aws_instance" "leofigyxvm" {
    ami = "${data.aws_ami.ubuntu_linux.id}"
    subnet_id = "${aws_default_subnet.default_az.id}"
    vpc_security_group_ids = ["${module.web_server_sg.this_security_group_id}", "${aws_security_group.admin_sg.id}"]
    instance_type = "t2.nano"
    key_name = "${var.ami_key_pair_name}"
    associate_public_ip_address=true
    tags = {
        Name ="leoX"
    }
}
