module.exports = {
    title: '王小右笔记',
    description: 'Just For Fun',
    lang: 'zh-CN',
    heroImage: "/logo.png",
    base: "/notes/",
    themeConfig: {
        repo: 'CoderCharm/notes',
        docsDir: 'docs',

        editLinks: true,
        editLinkText: 'Edit this page on GitHub',
        lastUpdated: 'Last Updated',

        nav: [
            {
                text: 'Golang',
                link: '/Golang/',
                activeMatch: '^/Golang/'
            },
            {
                text: 'Python',
                link: '/Python/',
                activeMatch: '^/Python/'
            },
            {
                text: 'JavaScript',
                link: '/JavaScript/',
                activeMatch: '^/JavaScript/'
            },
            {
                text: 'Linux',
                link: '/Linux/',
                activeMatch: '^/Linux/'
            },
            {
                text: 'Database',
                link: '/Database/',
                activeMatch: '^/Database/'
            },
        ],

        sidebar: {
            '/guide/': getGuideSidebar(),
            '/Golang/': getGolangConfig(),
            // '/': getGuideSidebar()
        }
    }
};

function getGuideSidebar() {
    return [
        {
            text: '笔记目录',
            children: [{text: '说明', link: '/guide/index'}]
        },
        {
            text: '分类',
            children: [
                {text: 'Golang', link: '/Golang/index'},
                {text: 'Python', link: '/Python/index'},
                {text: 'JavaScript', link: '/JavaScript/index'},
                {text: 'Linux', link: '/Linux/index'},
                {text: 'Database', link: '/Database/index'},
            ]
        }
    ]
}


function getGolangConfig() {
    return [
        {
            text: 'Go基础',
            children: [{text: '说明', link: '/Golang/index'}]
        },
        {
            text: 'Go进阶',
            children: [
                {
                    text: 'RabbitMQ',
                    children: [
                        {text: 'RabbitMQ介绍', link: '/Golang/rabbitMq/ch01_desc'},
                    ]
                },
            ]
        }
    ]

}