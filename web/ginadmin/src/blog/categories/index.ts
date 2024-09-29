import CategoryIcon from '@mui/icons-material/Category';

import CategoriesCreate from './CategoriesCreate';
import CategoriesEdit from './CategoriesEdit';
import CategoriesList from './CategoriesList';
import CategoriesShow from "./CategoriesShow"
export default {
    icon: CategoryIcon,
    create: CategoriesCreate,
    edit: CategoriesEdit,
    list: CategoriesList,
    show: CategoriesShow,
    recordRepresentation: 'reference',
}
