import TextField, { TextFieldProps } from "@mui/material/TextField";
export default function TextFields(props: TextFieldProps) {
  return <TextField type={props.type} placeholder="simple txt" />;
}
