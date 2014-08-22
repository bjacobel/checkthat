'use strict';

angular.module('CheckCMS')
    .controller('UsersCtrl', ['UsersService', '$scope', function (UsersService, $scope) {
        $scope.page = 'users';
        $scope.showUserCreateForm = false;

        $scope.newFirstName = null;
        $scope.newLastName = null;
        $scope.newTel = null;
        $scope.newNfcSerial = null;

        UsersService.getUserList(function(data){
            $scope.users = data;
        });

        $scope.toggleUserCreateForm = function(params){
            $scope.showUserCreateForm = !$scope.showUserCreateForm;
        }

        $scope.postUser = function(){
            if ( $scope.newFirstName == null || $scope.newLastName == null || $scope.newTel == null || $scope.newNfcSerial == null) {
                return false;
            } else {
                var params = {
                    "FirstName": $scope.newFirstName,
                    "LastName": $scope.newLastName,
                    "Tel": $scope.newTel,
                    "NfcSerial": $scope.newNfcSerial
                }
                UsersService.postUser(params, function(data){
                    console.log(data);
                });
            }

        };
    }]);
