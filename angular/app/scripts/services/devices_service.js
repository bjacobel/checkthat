'use strict';

angular.module('CheckCMS')
    .factory('DevicesService', ['Restangular',
        function DevicesService(Restangular) {
            return {
                getDeviceList: function(callback){
                    Restangular.one('devices').get().then(function(data){
                        callback(data);
                    });
                },
                getDevice: function(id, callback){
                    Restangular.one('devices', id).get().then(function(data){
                        callback(data);
                    });
                }
            };
        }
    ]);