output "id" {
    description = "Instance created id"
    value = "${aws_instance.leofigyxvm.id}"
}

output "public_dns" {
    description = "DNS Assigned"
    value = "${aws_instance.leofigyxvm.private_dns}"
}
output "public_ip" {
    description = "Public ip"
    value = "${aws_instance.leofigyxvm.public_ip}"
}
