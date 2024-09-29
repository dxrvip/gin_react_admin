import { Show, SimpleShowLayout, TextField ,useRecordContext} from 'react-admin';

const CategoriesShow = () => {

    const CategoryTitle = () => {
        const record = useRecordContext();
        return <span>查看: {record ? record.name : ''}</span>;
    }

    return (
    <Show title={<CategoryTitle />}>
        <SimpleShowLayout>
            <TextField source="id" />
            <TextField source="name" />
        </SimpleShowLayout>
    </Show>
    )
};


export default CategoriesShow;
