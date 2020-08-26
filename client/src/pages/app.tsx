import React from "react";

import { Dispatch, AnyAction, bindActionCreators } from "redux";
import { connect } from "react-redux";
// import { commonActions } from "../store/actions/common.actions";
import { IAppState } from "../store/reducers";

const mapStateToProps = (state: IAppState): IAppState => {
    return {
        commonState: state.commonState,
    };
};

const mapDispatchToProps = (dispatch: Dispatch<AnyAction>) => bindActionCreators(
    {
    },
    dispatch
);

type state = {};

class App extends React.Component<ReturnType<typeof mapStateToProps> & ReturnType<typeof mapDispatchToProps>, state> {
    constructor(props: any) {
        super(props);
        this.state = {
            value: ''
        };
    }



    handleChange = (event: any) => {
        this.setState({ value: event.target.value });
    }

    handleSubmit = (event: any) => {
        event.preventDefault();
    }

    render() {
        return (
            <div>
                App
            </div>
        )
    }
}

export default connect(mapStateToProps, mapDispatchToProps)(App);