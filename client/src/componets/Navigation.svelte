<script>
  import {link} from 'svelte-spa-router';
</script>
<header>
  <!-- Fixed navbar -->
  <nav class="navbar navbar-expand-md navbar-dark fixed-top bg-dark" style="background-color: #1f2937!important;">
    <div class="container-fluid">
      <a class="navbar-brand" href="/">CODE</a>
      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarCollapse" aria-controls="navbarCollapse" aria-expanded="false" aria-label="Toggle navigation">
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarCollapse">
        <ul class="navbar-nav me-auto mb-2 mb-md-0">
          <li class="nav-item">
            <a class="nav-link active" aria-current="page" use:link href="/">Home</a>
          </li>
          <li class="nav-item">
            <a class="nav-link" use:link href="/about">About</a>
          </li>
          <li class="nav-item dropdown">
            <a class="nav-link dropdown-toggle"  id="navbarDropdown"  role="button" data-bs-toggle="dropdown" aria-expanded="false">
              Master
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarDropdown">
              <li><a class="dropdown-item" use:link href="/company">Company</a></li>
              <li><a class="dropdown-item" use:link href="/pasaran">Pasaran</a></li>
            </ul>
          </li>
        </ul>
        <form class="d-flex">
          <button class="btn btn-warning" type="submit">Logout</button>
        </form>
      </div>
    </div>
  </nav>
</header>