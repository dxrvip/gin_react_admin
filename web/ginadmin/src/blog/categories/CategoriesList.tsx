import { List, Datagrid, TextField, DateField, BooleanField } from 'react-admin';

const CategoriesList = () => (
    <List resource='category'>
        <Datagrid>
            <TextField source="id" />
            <TextField source="name" />
  
        </Datagrid>
    </List>
);

export default CategoriesList;