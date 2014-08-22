'use strict';

/**
 * @ngdoc overview
 * @name CheckCMS
 * @description
 * # CheckCMS
 *
 * Main module of the application.
 */
 angular.module('CheckCMS', [
    'ngAnimate',
    'ngCookies',
    'ngRoute',
    'ngSanitize',
    'ngTouch',
    'restangular',
]).config(function ($routeProvider) {
    $routeProvider
    .when('/devices', {
        templateUrl: 'views/devices.html',
        controller: 'DevicesCtrl'
    })
    .when('/', {
        templateUrl: 'views/devices.html',
        controller: 'DevicesCtrl'
    })
    .when('/users', {
        templateUrl: 'views/users.html',
        controller: 'UsersCtrl'
    })
    .otherwise({
        redirectTo: '/'
    });
}).config(function(RestangularProvider) {
    RestangularProvider.setBaseUrl('http://checkthat.herokuapp.com/api/v1/');
});
