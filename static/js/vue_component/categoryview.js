import Topic from './topic.js';

export default {

	components: { Topic },

	data () {
		return {
			topics: [],
			categoryId: this.$route.params.categoryId,
			categoryName: this.$route.params.categoryName
		}
	},
	mounted () {
		console.log('api/categories/' + this.$route.params.categoryId + '/topics');
		axios.get('api/categories/' + this.$route.params.categoryId + '/topics').then((response) => {
			this.topics = response.data;
		})
	},
	template :`<div>
		<h2>Các chủ đề trong danh mục {{ categoryName }}</h2>
		<topic v-for="topic in topics" :topic="topic" :key="topic.id"></topic>
	</div>`

}
