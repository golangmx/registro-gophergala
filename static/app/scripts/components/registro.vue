<template>
	<div class="row">
		<div class="ten wide centered column">
			<div class="ui segment">
				<div class="ui blue two item pointing menu">
					<a class="active item" v-link="{ path: '/registro' }">Registro</a>
					<a class="item" v-link="{ path: '/equipos' }">Equipos</a>
				</div>
				<div class="ten column grid">
					<div class="row">
						<form class="ui form" @submit.stop.prevent="send" :class="{ 'loading': formLoading, 'success': formSuccess, 'error': formError }">
							<div class="ui info message">
								<div class="header">Nota</div>
								<p>Sólo es necesario que un miembro del equipo llene el registro.</p>
							</div>
							<h4 class="ui dividing header">Datos del Equipo</h4>
							<div class="two fields">
								<div class="field">
									<label>Nombre</label>
									<input type="text" v-model="equipo.nombre"/>
								</div>
								<div class="field">
									<label>Título del Proyecto</label>
									<input type="text" v-model="equipo.proyecto"/>
								</div>
							</div>
							<h4 class="ui dividing header">Tus Datos</h4>
							<div class="two fields">
								<div class="field">
									<label>Nombre(s)</label>
									<input type="text" v-model="usuario.nombres"/>
								</div>
								<div class="field">
									<label>Apellidos</label>
									<input type="text" v-model="usuario.apellidos"/>
								</div>
							</div>
							<div class="field">
								<label>Identificación Oficial</label>
								<div class="ui left action input">
										<div class="ui selection dropdown">
											<input name="tipo_id" type="hidden" v-model="usuario.tipo_id"/>
											<i class="dropdown icon"></i>
											<div class="default text">Tipo</div>
											<div class="menu">
												<div class="item" data-value="0">IFE</div>
												<div class="item" data-value="1">Pasaporte</div>
											</div>
										</div>
										<input type="text" pattern="[0-9 ]+" v-model="usuario.numero_id"/>
								</div>
							</div>
						<h4 class="ui dividing header">Integrantes</h4>
							<div class="ui message">
								<p>Registra a tus compañeros de equipo aquí.</p>
							</div>
							<div class="two fields">
								<div class="field">
									<label>Nombre(s)</label>
									<input type="text" v-model="miembro.nombres" @keydown.enter.stop.prevent="addMember"/>
								</div>
								<div class="field">
									<label>Apellidos</label>
									<input type="text" v-model="miembro.apellidos" @keydown.enter.stop.prevent="addMember"/>
								</div>
							</div>
							<a class="right floated ui green basic button" :class="{ 'disabled': addMemberEnabled }" @click.stop.prevent="addMember">
								<i class="add icon"></i>
								Añadir al equipo
							</a>
							<h4 class="ui dividing header">Vista Previa</h4>
							<div class="six wide ui teal padded segment">
								<h2 class="ui centered header">{{ equipo.nombre }}</h2>
								<h3 class="ui centered header">{{ equipo.proyecto }}</h3>
								<div class="ui divider" v-if="showDivider"></div>
								<div class="ui three cards">
									<div class="card" v-if="showUserCard">
										<div class="content">
											<i class="left floated user icon"></i>
											<div class="ui basic right floated disabled mini icon button">
												<i class="remove icon"></i>
											</div>
										</div>
										<div class="content">
											<p class="center aligned">{{ usuario.nombres }}</p>
											<p class="center aligned">{{ usuario.apellidos }}</p>
										</div>
									</div>
									<div class="card" v-for="m in miembros">
										<div class="content">
											<i class="left floated user icon"></i>
											<a class="ui basic right floated mini icon button" @click.stop.prevent="removeMember(m)">
												<i class="remove icon"></i>
											</a>
										</div>
										<div class="content">
											<p class="center aligned">{{ m.nombres }}</p>
											<p class="center aligned">{{ m.apellidos }}</p>
										</div>
									</div>
								</div>
							</div>
							<div class="ui success message">
								<p>{{ formMessage }}</p>
							</div>
							<div class="ui error message">
								<p>{{ formMessage }}</p>
								<p>Puedes reportar errores <a href="https://github.com/golangmx/registro-gophergala">aquí</a>.<p>
							</div>
							<button class="large ui primary button" type="submit" :disabled="submitDisabled">Enviar</button>
						</form>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	export default {
		data: function() {
			return {
				formLoading: false,
				formError: false,
				formSuccess: false,
				formMessage: "",
				equipo: {
					nombre: "",
					proyecto: ""
				},
				usuario: {
					nombres: "",
					apellidos: "",
					tipo_id: 0,
					numero_id: ""

				},
				miembro: {
					nombres: "",
					apellidos: ""
				},
				miembros: []
			}
		},
		computed: {
			addMemberEnabled: function() {
				return this.miembro.nombres == "" || this.miembro.apellidos == "";
			},
			showUserCard: function() {
				return this.usuario.nombres != "" || this.usuario.apellidos != "";
			},
			showDivider: function() {
				return ((this.equipo.nombre != "" || this.equipo.proyecto != "") &&
							(this.showUserCard || this.miembros.length > 0));
			},
			submitDisabled: function() {
				return this.usuario.nombres == "" || this.usuario.apellidos == "" ||
						this.usuario.id == "" || this.equipo.nombre == "" ||
						this.equipo.proyecto == "";
			}
		},
		methods: {
			addMember: function(e) {
				let {nombres, apellidos} = this.miembro;
				this.miembros = this.miembros.concat({
					nombres: nombres,
					apellidos: apellidos,
					tipo_id: 0,
					numero_id: ""
				});
				[this.miembro.nombres, this.miembro.apellidos] = ["", ""];
			},
			removeMember: function(m) {
				this.miembros = this.miembros.filter((x) => {
					return x.nombres != m.nombres && x.apellidos != m.apellidos;
				});
			},
			clearForm: function() {
				this.usuario.tipo_id = 0;
				this.miembros = [];
				[this.equipo.nombre, this.equipo.proyecto,
					this.usuario.nombres, this.usuario.apellidos,
					this.usuario.numero_id, this.miembro.nombres,
					this.miembro.apellidos] = "";
			},
			send: function(ev) {
				this.formMessage = "";
				this.formError = false;
				this.formSuccess = false;
				this.formLoading = true;
				let u = this.usuario;
				u.tipo_id = +u.tipo_id;
				let m = this.miembros;
				m = m.concat(u);
				let e = this.equipo;
				e.miembros = m;
				this.$http.post('/api/teams', e).then((res) => {
					this.formLoading = false;
					if (res.status != 201) {
						this.formMessage = "¡Oops! Recibimos una respuesta que no esperábamos.";
						this.formError = true;
						return;
					}
					this.formMessage = "¡Gracias! Tu registro fue procesado correctamente.";
					this.formSuccess = true;
					this.clearForm();
				}, (err) => {
					this.formLoading = false;
					this.formError = true;
					switch (err.status) {
						case 500:
							this.formMessage = "¡Oops! Algo extraño sucedió en el servidor (Status 500 Internal Server Error).";
							break;
						case 400:
							this.formMessage = "Hmm… La información enviada al servidor tiene un formato incorrecto. (Status 400 Bad Request).";
							break;
						default:
							this.formMessage = "¡Oops! Hubo un error: Status " + res.status + ".";
							break;
					}
				});
			}
		}
	}
</script>
