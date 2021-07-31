<script>
	import Router from 'svelte-spa-router';
	import {wrap} from 'svelte-spa-router/wrap'
	import { Row, Container, Col } from 'sveltestrap'
	import Navigation from './componets/Navigation.svelte'
	import Footer from './componets/Footer.svelte'
	import Home from './pages/Home.svelte'
	import NotFound from './pages/Notfound.svelte'

	const routes = {
		'/': wrap({
			component:Home
		}),
		'*': NotFound,
	}
</script>
<Navigation />
<main class="flex-shrink-0">
	<Container style="margin-top: 70px;">
		<Row>
			<Col>
				<Router routes={routes} />
				
			</Col>
		</Row>
		
	</Container>
</main>
<Footer />