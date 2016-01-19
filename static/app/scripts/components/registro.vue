<template>
	<div class="centered row">
		<h1 class="ui centered header">Gopher Gala México 2016</h1>
	</div>
	<div class="row">
		<div class="ten wide centered column">
			<div class="ui segment">
				<div class="ui blue two item pointing menu">
					<a class="active item" v-link="{ path: '/registro' }">Registro</a>
					<a class="item" v-link="{ path: '/equipos' }">Consulta</a>
				</div>
				<div class="ten column grid">
					<div class="row">
						<form class="ui form">
							<h4 class="ui dividing header">Datos del Equipo</h4>
							<div class="two fields">
								<div class="field">
									<label>Nombre</label>
									<input type="text"/>
								</div>
								<div class="field">
									<label>Título del Proyecto</label>
									<input type="text"/>
								</div>
							</div>
							<h4 class="ui dividing header">Tus Datos</h4>
							<h4 class="ui dividing header">Integrantes</h4>
						</form>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	export default {
		replace: false
	}
</script>
