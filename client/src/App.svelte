<script>
	import { Row, Container, Col } from 'sveltestrap'
	import Navigation from './componets/Navigation.svelte'
	import Footer from './componets/Footer.svelte'
	import Card from './componets/card.svelte'
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
			<Col xxl="3" xl="3" lg="6" md="6" sm="6" xs="6"  style="cursor:pointer;padding:5px;margin:0px;">
				<Card />
			</Col>
			<Col xxl="3" xl="3" lg="6" md="6" sm="6" xs="6"  style="cursor:pointer;padding:5px;margin:0px;">
				<Card />
			</Col>
			<Col xxl="3" xl="3" lg="6" md="6" sm="6" xs="6"  style="cursor:pointer;padding:5px;margin:0px;">
				<Card />
			</Col>
			<Col xxl="3" xl="3" lg="6" md="6" sm="6" xs="6"  style="cursor:pointer;padding:5px;margin:0px;">
				<Card />
			</Col>
		</Row>
		<center>
			<Loader cssstyle={'height: 100%;margin:100px 0px 100px 0px;'} />
		</center>
	</Container>
</main>
<Footer />