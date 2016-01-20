<template>
	<div class="row">
		<div class="ten wide centered column">
			<div class="ui segment">
				<div class="ui blue two item pointing menu">
					<a class="item" v-link="{ path: '/registro' }">Registro</a>
					<a class="active item" v-link="{ path: '/equipos' }">Equipos</a>
				</div>
				<div class="ten column grid">
					<div class="row"> <!-- search -->
					</div>
					<div class="row">
						<div class="ui error message" v-if="error">
							<p>{{ message }}</p>
						</div>
					</div>
					<div class="row">
						<div class="ui four cards">
							<equipo-card
								v-for="e in equipos"
								:equpo="e"
								:index="$index">
							</equpo-card>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	import EquipoCard from './equipo-card.vue';

	export default {
		data: function() {
			return {
				equipos: [],
				error: false,
				message: ""
			};
		},
		ready: function() {
			this.$http.get('/api/teams').then((res) => {
				let j = JSON.parse(res.responseText);
				this.equipos = j;
			}, (err) => {
				this.error = true;
				this.message = "¡Oops! Algo salió mal.";
			});
		},
		components: {
			'equipo-card': EquipoCard
		}
	}
</script>
