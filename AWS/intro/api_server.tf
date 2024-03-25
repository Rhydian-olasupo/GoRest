provider "aws" {
    profile = "default"
    region = "eu-central-1"
}

resource "aws_instance" "api_server" {
    ami = "ami-023adaba598e661ac"
    instance_type = "t2.micro"
}

