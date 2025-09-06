pipeline {
  agent { label 'dagger' }

  environment {
    DAGGER_MODULE = "github.com/jpadams/greetings-api/ci"
  }
  stages {
    stage("dagger") {
      steps {
        sh '''
            dagger call check
        '''
      }
    }
  }
}
