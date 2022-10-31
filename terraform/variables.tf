variable "region" {
  default = "us-east-1"
}

variable "aws_access_key" {
  default = ""
}

variable "aws_secret_key" {
  default = ""
}

variable "name" {
  default = "greetings"
}

variable "domain" {
  default = "kpenfound.io"
}

variable "fqdn" {
  default = "greetings.kpenfound.io"
}

variable "image" {
  default = "kylepenfound/greetings:latest"
}