<script>
	import Navigation from './componets/Navigation.svelte'
	import Footer from './componets/Footer.svelte'
	export let name;
</script>

<Navigation />
<main class="flex-shrink-0">
	<div class="container" style="margin-top: 70px;">
		<h1>Hello {name}!</h1>
		<h2>FIBER SVELTE</h2>
	</div>
</main>
<Footer />