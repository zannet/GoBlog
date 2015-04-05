app.controller("PostsController", ['$scope','Posts', function($scope,Posts){
	Posts.all(function(data){
		$scope.posts = data.posts;
		if(!$scope.$$phase) {
			$scope.$apply();
		}
	});
}]);

app.controller("PostEditController", ['$scope', 'Posts','$stateParams', function($scope, Posts, $stateParams){

	var originalPost = {};

	Posts.get($stateParams.slug, function(data){
		$scope.post = data.post;
		originalPost = jQuery.extend({},$scope.post);
		if(!$scope.$$phase) {
			$scope.$apply();
		}
	});

	$scope.save = function(){
		Posts.update($stateParams.slug, $scope.post, function(){
			originalPost = jQuery.extend({},$scope.post);
			$scope.canSave = false;
		});
	}

	$scope.$watchCollection('post', function(newVal,oldVal){
		$scope.canSave = !objEqual($scope.post, originalPost);
		console.log($scope.canSave);
	});

	var objEqual = function(a,b){
		var result = true;
		for(var i in a){
			if(a[i] != b[i]){
				result = false;
				console.log(i + " is not equal");
			}
		}
		return result;
	}

}]);