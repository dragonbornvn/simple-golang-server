
export default {
  name: 'App',
  template: `
  <div>
    <div class="container">
    	<div class="navbar">
    		<div class="container-fluid">
    			<a class="navbar-brand" href="#">My forum</a>
    			<ul class="nav navbar-nav">
    				<li class="active">
    					<a href="static/calculator-vue.html">Calculator</a>
    				</li>
    				<li>
    					<a href="static/job management.html">Job management</a>
    				</li>
    			</ul>
    		</div>
    	</div>
      <router-view></router-view>
    </div>
  </div>
  `,
};