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
          'web/assets/css/g0.css': 'assets/scss/g0.scss',
        }
      },
      build: {
        options: {
          style: 'compressed',
          cacheLocation: 'cache/sass-cache'
        },
        files: {
          'web/assets/css/g0.css': 'assets/scss/g0.scss',
        }
      }
    },
    watch: {
      files: ['assets/scss/**/*'],
      tasks: ['sass']
    }
  });

  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('default', ['sass:dist']);
  grunt.registerTask('build', ['sass:build']);
};
