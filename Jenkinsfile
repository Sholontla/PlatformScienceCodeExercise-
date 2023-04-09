pipeline {
    agent any
    
    stages {
        stage('Build') {
            steps {
                sh 'docker build -t myproject .'
            }
        }
        

        stage('Deploy Microservices') {
                steps {
                    script {
                        for (service in ["analytics", "logger", "stores", "analytics", "main_network"]) {
                            dir(service) {
                                sh 'docker-compose up -d'
                                sh 'docker-compose run --rm owasp zap-baseline.py -t http://localhost:8002'
                                sh 'docker-compose run --rm jmeter -n -t /jmeter/scripts/register_order.jmx -Jbase_url=http://localhost:8002/service/store/concurrent'
                            }
                        }
                    }
                }
            }
        }
    }
