### Blogs
https://aws.plainenglish.io/build-a-docker-image-and-publish-it-to-aws-ecr-using-github-actions-f20accd774c3

https://github.com/rahmanfadhil/gin-bookstore

https://medium.com/appgambit/part-1-running-docker-on-aws-ec2-cbcf0ec7c3f8



------
### Comandos utiles
docker build -t golang-api-test .

docker run -p 80:8081 -it golang-api-test 


chmod 0400 ec2GolangTest.pem
ssh -i ec2GolangTest.pem ec2-user@3.83.138.91

//instalar docker manual
$ sudo yum update -y
$ sudo amazon-linux-extras install docker
$ sudo service docker start
$ sudo usermod -a -G docker ec2-user

//instalalr docker en user data
#! /bin/sh
yum update -y
amazon-linux-extras install docker
service docker start
usermod -a -G docker ec2-user
chkconfig docker on

aws ecr describe-repositories--region us-east-1

eval $(aws ecr get-login --no-include-email | sed 's|https://||')
docker pull 263497198111.dkr.ecr.us-east-1.amazonaws.com/golang-api-test:1.3


aws dynamodb scan --table-name merchands_activities --region us-east-1
