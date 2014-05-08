module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    deploy: grunt.file.readJSON('.deploy.json'),
    sass: {
      test: {
        options: {
          style: 'expanded',
          cacheLocation: 'cache/sass-cache'
        },
        files: {
          'web/res/css/g0.css': 'res/scss/g0.scss',
        }
      },
      build: {
        options: {
          style: 'compressed',
          cacheLocation: 'cache/sass-cache'
        },
        files: {
          'web/res/css/g0.css': 'res/scss/g0.scss',
        }
      }
    },
    'sftp-deploy': {
      test: {
        auth: {
          host: '<%= deploy.test.host %>',
          port: '<%= deploy.test.port %>',
          authKey: 'test'
        },
        src: '<%= deploy.test.src %>',
        dest: '<%= deploy.test.dest %>',
        server_sep: '<%= deploy.test.sep %>'
      },
      live: {
        auth: {
          host: '<%= deploy.live.host %>',
          port: '<%= deploy.live.port %>',
          authKey: 'live'
        },
        src: '<%= deploy.live.src %>',
        dest: '<%= deploy.live.dest %>',
        server_sep: '<%= deploy.live.sep %>'
      }
    },
    shell: {
      build: {
        command: 'pub build'
      }
    },
    watch: {
      files: ['res/scss/**/*'],
      tasks: ['sass']
    }
  });

  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-sftp-deploy');
  grunt.loadNpmTasks('grunt-shell');

  grunt.registerTask('default', ['sass:test']);
  grunt.registerTask('build-test', ['sass:test', 'shell:build']);
  grunt.registerTask('build-live', ['sass:build', 'shell:build']);

  grunt.registerTask('deploy-test', ['build-test', 'sftp-deploy:test']);
  grunt.registerTask('deploy-live', ['build-live', 'sftp-deploy:live']);

};
