// /d:/playground/go-interpreter/index.ts

type StateUpdater<S> = (newState: S | ((prevState: S) => S)) => void;

class React {
    private static states: any[] = [];
    private static stateIndex = 0;

    static useState<S>(initialState: S): [S, StateUpdater<S>] {
        const currentIndex = React.stateIndex;
        React.states[currentIndex] = React.states[currentIndex] || initialState;

        const setState: StateUpdater<S> = (newState) => {
            if (typeof newState === 'function') {
                React.states[currentIndex] = (newState as (prevState: S) => S)(React.states[currentIndex]);
            } else {
                React.states[currentIndex] = newState;
            }
            React.render();
        };

        React.stateIndex++;
        return [React.states[currentIndex], setState];
    }

    static render() {
        React.stateIndex = 0;
        // Your render logic here
    }
}

export default React;