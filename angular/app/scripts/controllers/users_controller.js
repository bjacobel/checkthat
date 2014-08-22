'use strict';

angular.module('CheckCMS')
    .controller('UsersCtrl', ['UsersService', '$scope', function (UsersService, $scope) {
        $scope.page = 'users';

        UsersService.getUserList(function(data){
            $scope.users = data;
        });
    }]);
