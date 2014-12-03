'use strict';

module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    browserify: {
      options: {
        watch : true,
        keepAlive: true
      },
      app: {
        src: 'src/main.js',
        dest: 'build/frontend.js'
      }
    }
  });

  grunt.loadNpmTasks('grunt-browserify');

  // Default task(s).
  grunt.registerTask('default', ['browserify']);

};
