import EventEmitter from 'events';
import axios from 'axios';

const store = new EventEmitter();
let cache = null;
let cacheInvalidated = false;

export default store;

store.fetchEquipos = () => {
	return new Promise((resolve, reject) => {
		if (cache !== null && cacheInvalidated !== true) {
			resolve(cache);
		} else {
			axios.get('/api/teams')
				.then(res => {
					cache = res.data;
					cacheInvalidated = false;
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

store.postEquipo = e => {
	return new Promise((resolve, reject) => {
		axios.post('/api/teams', e).then(res => {
			if (res.status !== 201) {
				reject('¡Oops! Recibimos una respuesta que no esperábamos.');
			} else {
				cacheInvalidated = true;
				resolve('¡Gracias! Tu registro fue procesado correctamente.');
			}
		}).catch(err => {
			switch (err.status) {
				case 500:
					reject('¡Oops! Algo extraño sucedió en el servidor (Status 500 Internal Server Error).');
					break;
				case 400:
					reject('Hmm… La información enviada al servidor tiene un formato incorrecto. (Status 400 Bad Request).');
					break;
				default:
					reject('¡Oops! Hubo un error: Status ' + err.status + '.');
					break;
			}
		});
	});
};
