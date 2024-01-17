pipeline {
    agent any
    environment {
        // Ganti dengan informasi credential Docker Anda
        DOCKER_CREDENTIAL_ID = 'your-docker-credential-id'
    }
    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }
        stage('Build & Dockerize') {
            steps {
                script {
                    // Ganti 'jumatberkah/jenkins-test' dengan nama image yang Anda inginkan
                    def imageName = 'jumatberkah/jenkins-test'

                    // Build Docker image
                    def dockerImage = docker.build(imageName)

                    // Ganti 'main' dengan nama executable yang dihasilkan oleh build Go Anda
                    dockerImage.inside('-v $PWD:/app') {
                        sh 'go build -o main .'
                    }

                    // Simpan Docker image yang telah dibuat
                    dockerImage.save("build.tar")
                }
            }
        }
        stage('Push to Docker Hub') {
            steps {
                script {
                    // Load Docker image yang telah dibuat sebelumnya
                    def dockerImage = docker.load("build.tar")

                    // Push Docker image ke registry
                    docker.withRegistry('', DOCKER_CREDENTIAL_ID) {
                        dockerImage.push()
                    }
                }
            }
        }
    }
}