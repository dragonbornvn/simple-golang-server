
import Category from './category.js';
export default
{
	name:'Homeview',
	components: { Category },
	
	data() {
		return {
			categories:{}
		}
	},
	mounted () {
		axios.get('api/categories').then((response) => {
	        this.categories = response.data;
	    })
	},
	template :`<div>
		<category v-for="category in categories" :category="category" :key="category.id"></category>
	</div>`

}
