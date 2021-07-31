<script>
	import { Row, Container, Col } from 'sveltestrap'
	import Navigation from './componets/Navigation.svelte'
	import Footer from './componets/Footer.svelte'
	import Panel from './componets/panel.svelte'
	import Loader from './componets/Loader.svelte'
	
	const queryString = window.location.search
	const urlParams = new URLSearchParams(queryString)
	const token_browser = urlParams.get('token')
	let client_device = "WEBSITE"
</script>
<Navigation />
<main class="flex-shrink-0">
	<Container style="margin-top: 70px;">
		<Row>
			<Col>
				<Panel />
			</Col>
		</Row>
		
	</Container>
</main>
<Footer />