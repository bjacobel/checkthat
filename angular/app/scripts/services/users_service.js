'use strict';

angular.module('CheckCMS')
    .factory('UsersService', ['Restangular',
        function UsersService(Restangular) {
            return {
                getUserList: function(callback){
                    Restangular.all('users').getList().then(function(data){
                        callback(data);
                    });
                },
                getUser: function(id, callback){
                    Restangular.one('users', id).get().then(function(data){
                        callback(data);
                    });
                }
            };
        }
    ]);