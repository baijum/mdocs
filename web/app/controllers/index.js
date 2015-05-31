import Ember from 'ember';

export default Ember.Controller.extend({
  tocAgreed: false,
  actions: {
    signIn: function() {
      var self = this;
      Ember.$
	.ajax({
	  method: "POST",
	  url: "/api/v1/signin",
	  dataType: "json",
	  data: JSON.stringify({
	    username: self.get('username'),
	    password: self.get('password') })
	})
	.done(function( msg ) {
	  alert( "Data Saved: " + msg );
	});
    },
    signUp: function() {
      var self = this;
      Ember.$
	.ajax({
	  method: "POST",
	  url: "/api/v1/signup",
	  dataType: "json",
	  data: JSON.stringify({
	    user: {
	      fullname: self.get('fullname'),
	      username: self.get('username'),
	      email: self.get('email'),
	      password: self.get('password'),
	      toc: self.tocAgreed }})
	})
	.done(function( msg ) {
	  alert( "Data Saved: " + msg );
	});
    }
  }
});
