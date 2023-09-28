const app = Vue.createApp({
    // data, functions
    data() {
        return {
            url: 'https://google.com',
            showBooks: true,
            books: [
                { title: 'name of the wind', author: 'patrick rothfuss', img:'assets/1.jpg', isFav: true },
                { title: 'the way of the kings', author: 'brandon sanderson', img:'assets/2.jpg', isFav: false },
                { title: 'the final empire', author: 'brandon sanderson', img:'assets/3.jpg', isFav: true},
            ]
        }
    },
    methods: {
        toggleShowBooks() {
            this.showBooks = !this.showBooks;
        },
        setFav(book) {
            book.isFav = !book.isFav;
        }
    },
    computed: {
        filteredBooks() {
            return this.books.filter((book) => book.isFav);
        }
    }
});

app.mount('#app');

console.log('hello');