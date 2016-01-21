import Vue from 'vue';
import Router from 'vue-router';
import Registro from './components/registro.vue';
import Equipos from './components/equipos.vue';
import Equipo from './components/equipo.vue';

Vue.config.debug = true;

Vue.use(Router);

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
		component: Equipo
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
