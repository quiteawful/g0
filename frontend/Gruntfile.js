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
    watch: {
      files: ['res/scss/**/*'],
      tasks: ['sass']
    }
  });

  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');

  grunt.registerTask('default', ['sass:dist']);
  grunt.registerTask('build', ['sass:build']);
};
