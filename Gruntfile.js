function transformJSXES6(code) {
  return require('grunt-react').browserify(code, {harmony: true});
}

module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),
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

  // Load the plugin that provides the "uglify" task.
  grunt.loadNpmTasks('grunt-react');
  grunt.loadNpmTasks('grunt-browserify');

  // Default task(s).
  grunt.registerTask('default', ['browserify']);

};
