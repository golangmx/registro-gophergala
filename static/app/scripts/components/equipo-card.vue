<template>
	<a class="card" v-link="{ name: 'equipo', params: { id: id } }">
		<div class="content">
			<div class="header">{{ equipo.nombre }}</div>
			<div class="subheader">{{ equipo.proyecto }}</div>
		</div>
		<div class="extra content">
			<p><i class="user icon"></i> {{ memberCount }}</p>
		</div>
	</a>
</template>

<script>
	export default {
		props: ['equipo', 'id'],
		computed: {
			memberCount: function() {
				let l = this.equipo.miembros.length
				return l > 1 ? l + ' miembros' : l + ' miembro';
			}
		}
	}
</script>
