app.factory('Posts', ['$http',function($http){
	return {
		all: function(callback){
			$http({
				method: 'GET',
				url: "/api/posts"
			}).success(function(data){
				callback(data);
			});
		},
		get: function(slug,callback){
			$http({
				method: 'GET',
				url: "/api/post/"+slug
			}).success(function(data){
				callback(data);
			})
		},
		update: function(slug,post,callback){
			$http({
				method: 'PUT',
				url: "/api/post/"+slug,
				data: post
			}).success(function(data){
				callback();
			});
		}
	}
}]);