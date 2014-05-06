module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    sass: {
      dist: {
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
      build: {
        auth: {
          host: '188.226.132.29',
          port: 22,
          authKey: 'key1'
        },
        src: 'build',
        dest: '/var/www/slemgrim.com/public/g0',
        server_sep: '/'
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

  grunt.registerTask('default', ['sass:dist']);
  grunt.registerTask('build', ['sass:build']);
  grunt.registerTask('deploy', ['sass:build', 'sftp-deploy']);
};
