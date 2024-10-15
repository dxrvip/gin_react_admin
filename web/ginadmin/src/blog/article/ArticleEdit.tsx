import { RichTextInput } from "ra-input-rich-text"
import {
    Edit, SimpleForm, TextInput, TextField,
    ReferenceInput,
    SelectInput,
    ImageInput,
    ImageField,
    minLength,
    required,
    maxLength,
    useRecordContext
} from "react-admin"


const ArticleEdit = () => {
    const PostTitle = () => {
        const record = useRecordContext()
        return `编辑 ID: ${record?.id}的文章`
    }
    return (
        <Edit title={<PostTitle />}>
            <SimpleForm>

                <TextInput source="title" label="文章标题" validate={[required(), maxLength(100), minLength(2)]} />
                <TextInput source="desc" multiline={true} label="文章摘要" validate={[maxLength(200)]} />

                {/* 文章分类 */}
                <ReferenceInput source="cid" reference="category">
                    <SelectInput label="文章分类" optionText="name" validate={[required()]} />
                </ReferenceInput>


                <ImageInput source="picture" accept={{ 'image/*': ['.png', '.jpg'] }}>
                    <ImageField source="src" title="title" />
                </ImageInput>
                {/* 文章内容 */}
                <RichTextInput fullWidth source="content" label="文章内容" validate={[required(), minLength(10)]} />
            </SimpleForm>
        </Edit>
    )
}



export default ArticleEdit