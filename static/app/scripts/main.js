import Vue from 'vue';
import Router from 'vue-router';
import Resource from 'vue-resource';
import Registro from './components/registro.vue';
import Equipos from './components/equipos.vue';

/* eslint-disable no-undef */
$(document).ready(() => {
	$('div.dropdown').dropdown();
});
/* eslint-enable no-undef */

Vue.config.debug = true;

Vue.use(Router);
Vue.use(Resource);

let router = new Router();

router.map({
	'/registro': {
		component: Registro
	},
	'/equipos': {
		component: Equipos
	},
	'/equipos/:id': {
		name: 'equipo',
		component: Equipos
	}
});

router.beforeEach(() => {
	window.scrollTo(0, 0);
});

router.redirect({
	'/': '/registro'
});

let App = Vue.extend({});

router.start(App, '#app');
