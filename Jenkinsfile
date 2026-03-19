pipeline {
    agent any

    stages {
        stage('Build Docker Image') {
            steps {
                sh 'docker build -t miapp .'
            }
        }

        stage('Run Container') {
            steps {
                sh 'docker run -d -p 8081:8080 miapp'
            }
        }
    }
}
