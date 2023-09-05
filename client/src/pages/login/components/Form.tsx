import { BookOutlined } from "@mui/icons-material";
import {
  Alert,
  Box,
  Button,
  Card,
  CardActions,
  CardContent,
  CardMedia,
  Checkbox,
  CircularProgress,
  FormLabel,
  Grid,
  Snackbar,
  TextField,
} from "@mui/material";
import { FormEvent, useEffect, useState } from "react";
import { Controller, useForm } from "react-hook-form";
import { useMutation } from "react-query";
import { logIn } from "../api";

export default function Form() {
  const { control, getValues } = useForm();
  const [open, setOpen] = useState(false);
  const { mutate, isLoading, isError, error } = useMutation<
    any,
    {
      response: {
        data: {
          message: string;
        };
      };
    },
    any
  >(logIn);

  const onsubmit = async (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    mutate(getValues());
  };
  useEffect(() => {
    setOpen(isError);
  }, [isError]);
  return (
    <form onSubmit={(e) => onsubmit(e)}>
      {isError && (
        <Snackbar
          open={open}
          onClose={() => setOpen((prev) => !prev)}
          color="error"
          autoHideDuration={3000}
        >
          <Alert severity="error">{error.response.data.message}</Alert>
        </Snackbar>
      )}
      <Grid
        sx={{ width: "100%", border: "1px red solid", height: "100vh" }}
        container
        alignContent={"center"}
        justifyContent={"center"}
      >
        <Grid item md={6} sm={6}>
          <Card>
            <CardMedia sx={{ display: "flex", justifyContent: "center" }}>
              <BookOutlined sx={{ fontSize: 200, color: "gray" }} />
            </CardMedia>
            <CardContent
              sx={{ display: "flex", flexDirection: "column", gap: "20px" }}
            >
              <Controller
                name="email"
                control={control}
                defaultValue={""}
                render={(params) => {
                  return (
                    <TextField
                      onChange={params.field.onChange}
                      value={params.field.value}
                      label={"Email"}
                      name="email"
                    />
                  );
                }}
              />
              <Controller
                name="password"
                defaultValue={""}
                control={control}
                render={(params) => {
                  return (
                    <TextField
                      onChange={params.field.onChange}
                      value={params.field.value}
                      label={"Password"}
                      name="password"
                      type="password"
                    />
                  );
                }}
              />
              <Box>
                <Controller
                  name="rememberMe"
                  control={control}
                  defaultValue={true}
                  render={({ field: { onChange, value } }) => (
                    <>
                      <Checkbox
                        sx={{ alignSelf: "start" }}
                        value={value}
                        onChange={onChange}
                      />
                      <FormLabel>Remember me</FormLabel>
                    </>
                  )}
                />
              </Box>
            </CardContent>
            <CardActions>
              <Button
                variant="contained"
                color="primary"
                fullWidth
                type="submit"
              >
                {isLoading ? <CircularProgress /> : "Login"}
              </Button>
            </CardActions>
          </Card>
        </Grid>
      </Grid>
    </form>
  );
}
