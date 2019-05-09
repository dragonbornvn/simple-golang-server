export default {
    name:'Topic',
    props: ['topic'],
    template:`<div class="panel panel-success">
        <div class="panel-heading">
            <router-link :to="{ name: 'Topic', params: { topicId: topic.id } }" v-text="topic.title"></router-link>
        </div>
        <div class="panel-body">
            <span>by {{ topic.user.name }}</span>
            <span>{{ topic.time }}</span>
        </div>
        <div class="panel-footer">
            <span>{{ topic.views }}</span>
            <small>Lượt xem</small>
        </div>
    </div>  `
}
