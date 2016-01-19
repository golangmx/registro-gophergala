import Vue from 'vue';
import Router from 'vue-router';
import Registro from './components/registro.vue';

Vue.use(Router);

let router = new Router();

router.map({
	'/registro': {
		component: Registro
	},
	'/equipos': {
		//component: Equipos
	},
	'/equipos/:id': {
		//component: Equipo
	}
});

router.beforeEach(() => {
	window.scrollTo(0, 0);
});

router.redirect({
	'*': '/registro'
});

router.start(Registro, '#app');
