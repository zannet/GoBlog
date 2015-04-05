var app = angular.module('admin',['ui.router']);

app.config(function($stateProvider, $urlRouterProvider) {
  // For any unmatched url, redirect to /state1
  $urlRouterProvider.otherwise("/posts");
  //
  // Now set up the states
  $stateProvider
    .state('posts', {
      url: "/posts",
      templateUrl: "templates/posts.html",
      controller: "PostsController"
    })
    .state('post', {
      url: "/post/:slug",
      templateUrl: "templates/postedit.html",
      controller: "PostEditController"
    })
});