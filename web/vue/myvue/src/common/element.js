import Vue from 'vue'
import 'element-ui/lib/theme-chalk/index.css'
import {
    Button,
    Form,
    FormItem,
    Input,
    Message,
    Container,
    Header,
    Aside,
    Main,
    Col,
    Menu,
    Submenu,
    MenuItemGroup,
    MenuItem,
    Row,
   Breadcrumb,
   BreadcrumbItem,
   Card,
  Table,
  TableColumn,
  Tooltip,
  Pagination,
  Dialog,
  MessageBox,
  Badge,
  Step,
  Steps,
  Carousel,
  CarouselItem
} from 'element-ui'

// fade/zoom 等
import 'element-ui/lib/theme-chalk/base.css';
// collapse 展开折叠
import CollapseTransition from 'element-ui/lib/transitions/collapse-transition';
Vue.component(CollapseTransition.name, CollapseTransition)
Vue.prototype.$message = Message
Vue.prototype.$confirm = MessageBox.confirm
Vue.use(Button)
Vue.use(Form)
Vue.use(FormItem)
Vue.use(Input)
Vue.use(Container)
Vue.use(Header)
Vue.use(Aside)
Vue.use(Main)
Vue.use(Col)
Vue.use(Menu)
Vue.use(Submenu)
Vue.use(MenuItemGroup)
Vue.use(MenuItem)
Vue.use(Row)
Vue.use(Breadcrumb)
Vue.use(BreadcrumbItem)
Vue.use(Card)
Vue.use(Table)
Vue.use(TableColumn)
Vue.use(Tooltip)
Vue.use(Pagination)
Vue.use(Dialog)
Vue.use(Badge)
Vue.use(Step)
Vue.use(Steps)
Vue.use(Carousel)
Vue.use(CarouselItem)
