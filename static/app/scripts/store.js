import EventEmitter from 'events';
import axios from 'axios';

const store = new EventEmitter();
let cache = null;

export default store;

store.fetchEquipos = () => {
	return new Promise((resolve, reject) => {
		if (cache != null) {
			resolve(cache);
		} else {
			axios.get('/api/teams')
				.then(res => {
					cache = res.data;
					resolve(cache);
				}).catch(res => {
					reject('¡Oops! Algo salió mal. Status ' + res.status + '.');
				});
		}
	});
};

store.fetchEquipo = id => {
	return new Promise((resolve, reject) => {
		store.fetchEquipos()
			.then(equipos => {
				let n = equipos.filter(e => (e.id === id));
				if (n.length !== 1) {
					reject('¡Oops! Este equipo no existe.');
				} else {
					resolve(n[0]);
				}
			}).catch(err => {
				reject(err);
			});
	});
};
