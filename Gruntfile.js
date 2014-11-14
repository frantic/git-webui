function transformJSXES6(code) {
  return require('grunt-react').browserify(code, {harmony: true});
}

module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
    watch: {
      files: ['src/*.js'],
      tasks: ['browserify'],
    },
    browserify: {
      options: {
        transform: [ transformJSXES6 ]
      },
      app: {
        src: 'src/main.js',
        dest: 'build/frontend.js'
      }
    }
  });

  grunt.loadNpmTasks('grunt-react');
  grunt.loadNpmTasks('grunt-browserify');
  grunt.loadNpmTasks('grunt-contrib-watch');

  // Default task(s).
  grunt.registerTask('default', ['browserify']);

};
