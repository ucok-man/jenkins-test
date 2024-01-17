pipeline {
    agent any
    environment {
        // Definisi variabel lingkungan yang diperlukan
        DOCKER_REGISTRY_CREDENTIALS = '94e58d8b-6783-4262-b1a0-96250f45f61c'
        DOCKER_IMAGE_NAME = 'jumatberkah/jenkins-test'
    }
    stages {
        stage('Checkout'){
            steps {
                checkout scm
            }
        }
        stage('Build'){
            steps {
                script {
                    // Ganti perintah ini sesuai dengan perintah yang diperlukan untuk membangun aplikasi Go Anda
                    sh 'go build -o myapp'
                }
            }
        }
        stage('Dockerize'){
            steps {
                script {
                    // Ganti perintah ini sesuai dengan perintah yang diperlukan untuk membuat Docker image
                    sh 'docker build -t $DOCKER_IMAGE_NAME .'
                }
            }
        }
        stage('Push to Docker Hub'){
            steps {
                script {
                    // Push Docker image ke Docker Hub
                    docker.withRegistry('', DOCKER_REGISTRY_CREDENTIALS) {
                        docker.image(DOCKER_IMAGE_NAME).push()
                    }
                }
            }
        }
    }
}