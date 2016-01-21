<template>
	<div class="row">
		<div class="ten wide centered column">
			<div class="ui segment">
				<!--<div class="ui blue two item pointing menu">
					<a class="item" v-link="{ path: '/registro' }">Registro</a>
					<a class="active item" v-link="{ path: '/equipos' }">Equipos</a>
				</div>-->
				<div class="ui secondary segment">
					<div class="ui large breadcrumb">
						<a class="section" v-link="{ path: '/equipos' }">Equipos</a>
						<i class="right chevron icon divider"></i>
						<div class="active section">{{ nombre }}</div>
					</div>
				</div>
				<div class="ui divider"></div>
				<div class="ten column grid">
					<div class="row">
						<h2 class="ui centered header">{{ nombre }}</h2>
						<h3 class="ui centered sub header">{{ proyecto }}</h3>
					</div>
					<div class="ui horizontal divider">
						<i class="users icon"></i>
					</div>
					<div class="row">
						<div class="ui three cards">
							<div class="card" v-for="m in miembros">
								<div class="content">
									<i class="left floated user icon"></i>
								</div>
								<div class="content">
									<p class="center aligned">{{ m.nombres }}</p>
									<p class="center aligned">{{ m.apellidos }}</p>
								</div>
								<div class="extra content" v-if="m.numero_id.length > 0">
									<p class="meta">{{ m.tipo_id === 0 ? 'IFE' : 'Pasaporte' }} {{ m.numero_id }}</p>
								</div>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	import store from '../store.js';

	export default {
		data: function() {
			return {
				nombre: '',
				proyecto: '',
				miembros: []
			};
		},
		route: {
			data ({to}) {
				const id = +to.params.id;
				return store.fetchEquipo(id)
					.then(e => ({
						nombre: e.nombre,
						proyecto: e.proyecto,
						miembros: e.miembros
					})).catch(err => {
						console.log(err);
						return {
							nombre: '',
							proyecto: '',
							miembros: []
						};
					});
			}
		}
	}
</script>
