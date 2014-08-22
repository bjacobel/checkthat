'use strict';

angular.module('CheckCMS')
    .controller('DevicesCtrl', ['DevicesService', '$scope', function (DevicesService, $scope) {
        $scope.page = 'devices';

        DevicesService.getDeviceList(function(data){
            $scope.devices = data.CheckedOut.concat(data.CheckedIn);

            $scope.devices.map(function(curr, index, array){
                if (data.CheckedOut.indexOf(curr) != -1){
                    curr.status = "out";
                } else if (data.CheckedIn.indexOf(curr) != -1){
                    curr.status = "in";
                }
                return curr;
            })
        });
    }]);
