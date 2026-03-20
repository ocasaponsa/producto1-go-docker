pipeline {
    agent any

    stages {
        stage('Check files') {
            steps {
                sh 'ls -la'
                sh 'test -f Dockerfile'
                sh 'test -f main.go'
                sh 'test -f go.mod'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t miapp .'
            }
        }
    }
}
